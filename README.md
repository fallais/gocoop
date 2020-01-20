# GoCoop

![Coop](https://github.com/fallais/gocoop/blob/master/assets/coop.png)

**GoCoop** is a tool written in **Go** and **Angular** that helps you to manage your chicken coop !

## Objectives

The main objective is to protect the chickens against the **hungry fox** or the **greedy weasel**. To do so, we need to automaticaly open and close the door of the chicken coop, with two options :

- At a fixed time (for example *08h30*)
- Based on the sunset and sunrise (for example *30min after sunrise*)

> If you worry about using the sun based condition, be sure that the chickens always go to sleep at sunset. As the sentence says : **go to bed with the chickens**.

## Components

### The motor

I use the `Nextrox 37mm 12V 15RPM`. I chose this motor because of its torque : **250 N*cm**

![Nextrox](https://github.com/fallais/gocoop/blob/master/assets/nextrox.jpg)

### The motor driver

I use the `L293D`. It is capable of handling two motors.

![L293D](https://github.com/fallais/gocoop/blob/master/assets/L293D.jpg)

### The GPIO pins

Here are the GPIO pins that are used :

- 23 : connected to the **Input 1** of the **L293D**
- 24 : connected to the **Input 2** of the **L293D**
- 25 : connected to the **Enable 1** of the **L293D**

The GPIO of my **Raspberry 3 B+** are located as follow :

![GPIO](https://github.com/fallais/gocoop/blob/master/assets/gpio.png)

## Is it tested ?

Sure, I tried to do my best to add package tests because chickens deserv the best ! Moreover, I have been using it for more than one year at home, it has never failed since the begining.

## Interface

It also comes with an interface built with **Angular** to manage the coop.

Protected by a login.  
![Login](https://github.com/fallais/gocoop/blob/master/assets/login.png)

With a dashboard.  
![dashboard](https://github.com/fallais/gocoop/blob/master/assets/dashboard.png)

## Installation

First, I installed **raspbian-lite** on the Raspberry.  
Then, I updated all the packages.  
Finally, I installed **docker** with the convenience script.  
And, a basic **logrotate** configuration.

The Raspberry is ready to run the Docker container.

## Usage

### Docker

Deploy with a `docker-compose`.

```yaml
version: "3"
services:
  redis:
    image: redis
    container_image: redis
    restart: always
    networks:
      main:
        aliases:
          - redis

  gocoop:
    image: fallais/gocoop
    container_image: gocoop
    restart: always
    volumes:
      - ./config.yml:/usr/bin/config.yml
    ports:
      - 80:2015
    networks:
      main:
        aliases:
          - gocoop

networks:
  main:
    driver: bridge
```

#### Parameters

- config.yml : mandtory
- port : mandatory

### Configuration

The configuration file must be as follow :

```yaml
general:
  gui_username: admin
  gui_password: admin
  private_key: myK3yIsAwesome!
  redis_host: localhost:6379
  redis_password: 
coop:
  latitude: 42.525776
  longitude: 2.327727
  opening:
    mode: "time_based"
    value: "08h00"
  closing:
    mode: "sun_based"
    value: "-30m"
door:
  openening_duration: "65s"
  closing_duration: "60s"
```

#### Modes and values

Two modes are available :

- Time based (fixed time) : `time_based`
  - Value must be something like **HHhMM** : `08h00`
- Sun based (based on the sunrise and sunset) : `sun_based`
  - Value must be a valid Golang duration : `45m`

## Production ready

It is actually also used by a friend who have **160 chickens**. Below an overview of how it looks like.

![Door](https://github.com/fallais/gocoop/blob/master/assets/door.jpg)

## Licence

I do not set any licence as of now. Please do not use code for commercial purpose, instead, contact me and we could work together. Chickens are against money problems..