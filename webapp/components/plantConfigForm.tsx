import { useState } from "react";
import { ClientType, ConfigType } from "utils/types";
import { Component } from "react";
import { setConfig as postConfig } from "utils/req";

const Input = ({
  title,
  value,
  onChange
}: {
  title: string;
  value: string | number;
  onChange: (string) => void;
}) => (
  <div>
    <div>{title}:</div>
    <input
      value={value || ""}
      onChange={input => onChange(input.target.value)}
    />
  </div>
);

const Status = ({ active }: { active: boolean }) => (
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

const PlantConfigForm = ({
  client,
  onPost
}: {
  client: ClientType;
  onPost: (id: string, config: ConfigType) => void;
}) => {
  const { id, active, config: currentConfig } = client;
  const [config, setConfig] = useState(
    currentConfig || { plant: undefined, water: undefined }
  );
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
