## State

* If the dom re renders an input then the data inside of it is lost unless it is stored in state.
* ex:
``` jsx
<input
    type="text"
    value={this.state.name}
    onChange={(e) => this.setState({name: e.target.value})}
/>
```

