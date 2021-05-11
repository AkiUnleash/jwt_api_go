import * as React from 'react';

type props = {
  type: string,
  placeholder: string,
  id: string,
  value: string,
  setValue: React.Dispatch<React.SetStateAction<string>>
  label: string
}

const Textfield: React.FC<props> = (props: props) => {
  return (
    <div>
      <label>{props.label}</label>
      <input type={props.type} placeholder={props.placeholder}
        id={props.id}
        defaultValue={props.value}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          props.setValue(e.target.value)
        }}
      />
    </div>
  );
};
export default Textfield;