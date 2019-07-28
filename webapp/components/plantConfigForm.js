import { Component } from "react";
import { BASE_URL } from "utils/config";

const postConfig = (id, config) => {
  fetch(`${BASE_URL}/api/config`, {
    method: "POST",
    body: JSON.stringify({ ...config, id })
  });
};

class PlantConfigForm extends Component {
  state = {};
  render() {
    const { id } = this.props;
    return (
      <div key={id}>
        <div>Id: {id}</div>
        <input
          value={this.state["plant"] || ""}
          onChange={input => {
            const value = input.target.value;
            this.setState({
              plant: value !== "" ? value : undefined
            });
          }}
        />
        <input
          value={this.state["water"] || ""}
          onChange={input => {
            const value = parseInt(input.target.value);
            this.setState({
              water: value || undefined
            });
          }}
        />
        <button onClick={() => postConfig(id, this.state)}>Configure</button>
      </div>
    );
  }
}

export default PlantConfigForm;
