# GoCoop

![Coop](https://github.com/fallais/gocoop/blob/master/assets/coop.png)

**GoCoop** is a tool written in **Go** that helps you to manage your chicken coop !

## Objectives

The main objective is to protect the chickens against the hungry **fox** or the **greedy weasel**. To do so, we need to automaticaly open and close the door of the chicken coop, with two options :

- Fixed time (for example *08h30*)
- Regarding the sunset and sunrise (for example *30min after sunrise*)

> The chickens always go to sleep at sunset. You know the sentence : **go to bed with the chickens**.

## Components

### The motor

I use the `Nextrox 37mm 12V 15RPM`. I chose this motor because of its torque : **250 N*cm**

### The motor driver

I use the `L293D`. It is capable of handling two motors.

![L293D](https://github.com/fallais/gocoop/blob/master/assets/L293D.jpg)

### The GPIO pins

Here are the GPIO pins that are used :

- 23 : Input 1 (L293D)
- 24 : Input 2 (L293D)
- 25 : Enable 1 (L293D)

## Is it tested ?

Sure, chickens deserv the best ! The fox is **mercyless** !

## Usage

### Docker

Pull the image : `docker pull fallais/gocoop`
Run the image : ``

### Configuration

The configuration file must be as follow :

```yaml
latitude: 42.525776
longitude: 2.327727
gui_username: admin
gui_password: admin
static_dir: /app/frontend
opening:
  mode: "time_based"
  value: "08h00"
closing:
  mode: "sun_based"
  value: "-30m"
```

#### Modes and values

Two modes are available :

- Time based (fixed time) : `time_based`
  - Value must be something like **HHhMM** : `08h00`
- Sun based (based on the sunrise and sunset) : `sun_based`
  - Value must be a valid Golang duration : `45m`
