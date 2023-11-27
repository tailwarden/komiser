import { ChangeEvent } from 'react';

type CheckboxProps = {
  id?: string;
  checked?: boolean;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
};

function Checkbox({ id, checked, onChange }: CheckboxProps) {
  return (
    <input
      id={id}
      checked={checked}
      onChange={onChange}
      type="checkbox"
      className="grid h-4 w-4 appearance-none place-content-center rounded border-2 border-gray-300 bg-transparent before:h-[1rem] before:w-[1rem] before:origin-bottom-left before:scale-0 before:shadow-[inset_1rem_1rem_#fff] before:content-[''] before:[clip-path:polygon(28%_38%,41%_53%,75%_24%,86%_38%,40%_78%,15%_50%)] checked:border-darkcyan-500 checked:bg-darkcyan-500 checked:before:scale-100 hover:border-gray-400 checked:hover:border-transparent checked:hover:bg-darkcyan-700"
    />
  );
}

export default Checkbox;
