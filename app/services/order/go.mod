module github.com/nndergunov/deliveryApp/app/services/order

go 1.18

replace (
	github.com/nndergunov/deliveryApp/app/pkg/messagebroker  => C:/Users/Mykyta_Derhunov/Desktop/Code/deliveryApp/app/pkg/messagebroker
	github.com/nndergunov/deliveryApp/app/pkg/configreader  => C:/Users/Mykyta_Derhunov/Desktop/Code/deliveryApp/app/pkg/configreader
)

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.6
	github.com/nndergunov/deliveryApp/app/pkg/api v0.0.0-20220526132843-19b4d32b0c4f
	github.com/nndergunov/deliveryApp/app/pkg/configreader v0.0.0-20220526132843-19b4d32b0c4f
	github.com/nndergunov/deliveryApp/app/pkg/logger v0.0.0-20220526132843-19b4d32b0c4f
	github.com/nndergunov/deliveryApp/app/pkg/messagebroker v0.0.0-20220526132843-19b4d32b0c4f
	github.com/nndergunov/deliveryApp/app/pkg/server v0.0.0-20220526132843-19b4d32b0c4f
	github.com/volatiletech/sqlboiler/v4 v4.11.0
	github.com/volatiletech/strmangle v0.0.4
)

require (
	github.com/ericlagergren/decimal v0.0.0-20181231230500-73749d4874d5 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/rabbitmq/amqp091-go v1.3.4 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.11.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null/v8 v8.1.2 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220517211312-f3a8303e98df // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
