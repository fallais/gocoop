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

The GPIO of the **Raspberry 3** are located as follow :

![GPIO](https://github.com/fallais/gocoop/blob/master/assets/gpio-1.jpg)

## Is it tested ?

Sure, chickens deserv the best ! The fox is **mercyless** ! I have been using it for more than one year at home, it has never failed since.

## Usage

### Docker

Build the image for backend : `docker build -t gocoop -f build/docker/backend.Dockerfile`.  
Build the image for frontend : `docker build -t gocoop-frontend -f build/docker/frontend.Dockerfile`.  
Deploy with a `docker-compose`.

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
