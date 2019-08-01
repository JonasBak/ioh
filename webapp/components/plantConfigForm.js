import { Component } from "react";
import { setConfig } from "utils/req";
import { BASE_URL } from "utils/config";

class PlantConfigForm extends Component {
  constructor(props) {
    super(props);
    const { plant, water } = props.config || {};
    this.state = { plant, water };
  }
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
        <button
          onClick={() =>
            setConfig(id, this.state["plant"], this.state["water"])
          }
        >
          Configure
        </button>
      </div>
    );
  }
}

export default PlantConfigForm;
