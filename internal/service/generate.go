package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i UserService -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i UserCacheService -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i UserConsumerService -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i UserProducerService -o ./mocks/ -s "_minimock.go"
