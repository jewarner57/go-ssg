## OOP Inheritance

### Inheritance In JS
``` javascript

class Sprite() {
  constructor(x, y, color) {
    this.x = x;
    this.y = y;
    this.color = color;
  }
}

class Ball extends Sprite {
  constructor(x, y, color, radius) {

    // super calls the parent class's constructor
    super(x, y, color);

    this.radius = radius;
  }
}

```

### REPL - https://repl.it/@JonathanWarner2/top-canvas-inheritance