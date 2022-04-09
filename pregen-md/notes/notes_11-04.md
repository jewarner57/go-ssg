## FEW Notes - Javascript OOP

### OOP examples

class Brick {
    constructor(x, y) {
        this.x = x
        this.y = y
        this.status = true
    }

    break() {
        this.status = false
    }

    display() {
        //display the block
    }
}

brick = new Brick(100, 100)