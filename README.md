# Beresheet lunar lander

This simulation was built by inspiration of SpaceIL project - Beresheet. Which is mission to land on the moon, more details can be found -  https://en.wikipedia.org/wiki/SpaceIL

Beresheet spacecraft unfortunately didn't manage to land properly on the moon, there are many speculations and thoughts about the real reason of this failure, alot of reasons can be found on the interenet or SpaceIL site.

The main goal of this simulator is trying to make landing process successfully which will do it fully automatically using all sensors available in the craft.

We have here few parameters we count on: 
- Vertical Speed, Horizontal Speed, Distance, Angle-rotation, Altitude, Acceleration Rate, Weight, Fuel.

All these parameters and few more will calculate every frame; it can be either 1 per second or even more iterations per second.

There is no graphics in this simulation, so the only way to see what happened is debug mode on which prints details every frame.
More info about the Beresheet craft can be found -  https://en.wikipedia.org/wiki/Beresheet

## How to run?
### v1.0.0
In class, we demonstrated this version. There are two algorithms in this version: BoazLanding (leaves 16 liters of fuel) and Landing (leaves 21 liters of fuel).
Run the simulation by entering the `make run` command.
### v2.0.0
Following our demonstration of algorithms to students in the classroom, we designed a new one.
- `make runBoaz` - runs the algorithm that we received along with the assignment (leaves 16 liters of fuel).
- `make run` - runs the algorithm that we showed in the class (leaves 21 liters of fuel).
- `make runTwoPID` - runs the new algorithm (leaves 24 liters of fuel).
