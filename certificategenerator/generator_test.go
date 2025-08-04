package certificategenerator_test

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/certificategenerator"
)

const Cert_Dir = "certificate"

func TestCertificateGenerator(t *testing.T) {
	t.Run("Создание приватного ключа", func(t *testing.T) {
		privateKeyFile, err := os.Create(path.Join(Cert_Dir, "private_key.pem"))
		if err != nil {
			log.Fatalln(err)
		}

		privateKey, err := certificategenerator.GeneratePrivateKey()
		if err != nil {
			log.Fatalln(err)
		}

		privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			log.Fatalln(err)
		}

		err = pem.Encode(privateKeyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes})
		if err != nil {
			log.Fatalln(err)
		}
	})

	t.Run("Создание корневого сертификата (Root Certificate)", func(t *testing.T) {
		certFile, err := os.Create(path.Join(Cert_Dir, "root_certificate.pem"))
		if err != nil {
			log.Fatalln(err)
		}

		certBytes, err := certificategenerator.GenerateCertificate()
		if err != nil {
			log.Fatalln(err)
		}

		err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
		if err != nil {
			log.Fatalln(err)
		}

		//Теперь, когда у нас есть новый корневой центр сертификации, мы можем создать
		// промежуточные центры сертификации. Это означает, что наш корневой сертификат
		// может выдавать новые сертификаты центра сертификации для промежуточных центров
		// сертификации, которые, в свою очередь, могут выдавать сертификаты для конечных
		// объектов, создавая цепочку доверия.
		//Используя сертификат и закрытый ключ корневого центра сертификации, мы можем
		// создать и подписать новый промежуточный сертификат.

		//Посмотреть подробную инфрмацию о сертификате можно командой:
		// openssl x509 -in intermediate_ca_certificate.pem -text -noout
	})
}
