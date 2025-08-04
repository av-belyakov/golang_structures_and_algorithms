package certificategenerator

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"time"
)

// GenerateCertificate генератор сертификата
func GenerateCertificate() (certBytes []byte, err error) {
	var serialNumber *big.Int

	//генерируем случайный серийный номер
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return
	}

	//создание временного сертификата
	template := x509.Certificate{
		SerialNumber: serialNumber,                   //уникальный номер сертификата
		NotBefore:    time.Now(),                     //дата начала действия сертификата
		NotAfter:     time.Now().Add(time.Hour * 24), //дата окончания действия сертификата
		KeyUsage:     x509.KeyUsageCertSign,          //открытый ключ в этом сертификате можно
		// использовать для проверки подписей в других сертификатах
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, //дополнительная
		// цель для конкретного приложения (в нашем случае это аутентификация на веб-сервере)
		Subject: pkix.Name{
			CommonName: "domain-example",
		}, //отличительное имя объекта
		//ниже это альтернативные имена субъекта (Subject Alternative Names, SAN), которые
		// определяют личность владельца сертификата. Это могут быть DNS-имена (чаще всего
		// используются для доменных имён), IP-адреса, адреса электронной почты или URI.
		// Для локального тестирования мы создаём сертификат для домена localhost и обратных
		// IPv4- и IPv6-адресов.
		IPAddresses:           []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:              []string{"domain-example.net", "domain-example.ru"},
		EmailAddresses:        []string{"domain-example@email.net", "domain-example@email.ru"},
		Version:               3,    //версия сертификата
		BasicConstraintsValid: true, //флаг, который указывает на то, что ограничения
		// действительны, а не просто имеют нулевое значение
		IsCA: true, //это сертификат центра сертификации
	}

	parent := template
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		return
	}

	//создание самого сертификата
	certBytes, err = x509.CreateCertificate(
		rand.Reader,
		&template,             // временный сертификат
		&parent,               // родительский сертификат
		&privateKey.PublicKey, // открый ключ заявителя на получение сертификата
		privateKey,            // закрытый ключ лица, подписывающего сертификат
	)
	if err != nil {
		return
	}

	return
}

// GeneratePrivateKey новый приватный ключ
func GeneratePrivateKey() (privateKey *ecdsa.PrivateKey, err error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}
