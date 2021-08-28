provider "aws" {
    region = "us-east-1"
}

//Ищем последний дистрибутив Ubuntu 20.04
data "aws_ami" "ubuntu" {
    most_recent = true
    filter {
        name   = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
    }
    filter {
        name = "virtualization-type"
        values = ["hvm"]
    }
    owners = ["099720109477"]   // Canonical
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

// Выделяем диспазаон адресов для проекта
resource "aws_vpc" "netology_vpc" {
  cidr_block = "172.16.0.0/16"
}

//Из диапазона выделям отдельную подсеть для сервера
resource "aws_subnet" "subnet_for_test_servers" {
  vpc_id            = aws_vpc.netology_vpc.id
  cidr_block        = "172.16.1.0/24"
  availability_zone = "us-east-1"
}

//Создаем виртуальный интерфейс и привязываем на него IP адрес
resource "aws_network_interface" "test_server" {
  subnet_id   = aws_subnet.subnet_for_test_servers.id
  private_ips = ["172.16.1.2"]
}

//Описываем создание тестового сервера
resource "aws_instance" "netology_test" {
    ami = data.aws_ami.ubuntu.id   
    instance_type = "t3.micro"

//Описание интерфейса
    network_interface {
        network_interface_id = aws_network_interface.test_server.id
        device_index         = 0
    }

    tags = {
        Name = "netology-test"
    }

}
