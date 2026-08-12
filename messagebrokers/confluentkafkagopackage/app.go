package confluentkafkagopackage

import "errors"

// New настраивает модуль взаимодействия с API Kafka
func New(counter Counter, logger Logger, opts ...KafkaApiOptions) (*KafkaApiModule, error) {
	api := &KafkaApiModule{
		counter: counter,
		logger:  logger,
		settings: kafkaApiSettings{
			cachettl: 15,
		},
		chFromModule: make(chan ChOutputSettings),
		chToModule:   make(chan ChInputSettings),
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	return api, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		n.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		n.settings.port = v

		return nil
	}
}

// WithCacheTTL время жизни для кэша хранящего функции-обработчики запросов к модулю
func WithCacheTTL(v int) KafkaApiOptions {
	return func(th *KafkaApiModule) error {
		if v <= 10 || v > 86400 {
			return errors.New("the lifetime of a cache entry should be between 10 and 86400 seconds")
		}

		th.settings.cachettl = v

		return nil
	}
}

// WithNameRegionalObject наименование
func WithNameRegionalObject(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.nameRegionalObject = v

		return nil
	}
}

// WithTopicsSubscription 'слушатель' разных топиков
func WithTopicsSubscription(v map[string]string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'topics' cannot be empty")
		}

		n.topics = v

		return nil
	}
}

// WithLocationCertificateCA местоположение корневого сертификата
func WithLocationCertificateCA(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.locationCertificateCA = v

		return nil
	}
}

// WithLocationClientCertificate местоположение клиентского сертификата в формате PEM
func WithLocationClientCertificate(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.locationClientCertificate = v

		return nil
	}
}

// WithLocationClientKey местоположение приватного клиентского ключа в формате PEM
func WithLocationClientKey(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.locationClientKey = v

		return nil
	}
}
