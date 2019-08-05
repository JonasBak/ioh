import { useState } from "react";
import { Component } from "react";
import { setConfig as postConfig } from "utils/req";

const Input = ({ title, value, onChange }) => (
  <div>
    <div>{title}:</div>
    <input
      value={value || ""}
      onChange={input => onChange(input.target.value)}
    />
  </div>
);

const Status = ({ active }) => (
  <div>
    <div
      style={{
        backgroundColor: active ? "green" : "red",
        width: "1em",
        height: "1em",
        borderRadius: "0.5em",
        display: "inline-block"
      }}
    />
    {active ? "Online" : "Offline"}
  </div>
);

const PlantConfigForm = ({ client, onPost }) => {
  const { id, active, config: currentConfig } = client;
  const [config, setConfig] = useState(currentConfig || {});
  return (
    <div>
      <div>Id: {id}</div>
      <Status active={active} />
      <Input
        title="Plant"
        value={config.plant}
        onChange={value =>
          setConfig({
            ...config,
            plant: value !== "" ? value : undefined
          })
        }
      />
      <Input
        title="Water"
        value={config.water}
        onChange={value =>
          setConfig({
            ...config,
            water: parseInt(value) || undefined
          })
        }
      />
      <button
        onClick={async () => {
          const result = await postConfig(id, config.plant, config.water);
          onPost(id, result.data.setConfig);
        }}
      >
        Configure
      </button>
    </div>
  );
};

export default PlantConfigForm;
