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
      className="appearance-none bg-transparent border-2 w-4 h-4 border-gray-300 hover:border-gray-400 rounded grid place-content-center before:content-[''] before:w-[1rem] before:h-[1rem] before:[clip-path:polygon(28%_38%,41%_53%,75%_24%,86%_38%,40%_78%,15%_50%)] before:scale-0 before:origin-bottom-left before:shadow-[inset_1rem_1rem_#fff] checked:before:scale-100 checked:bg-komiser-600 checked:border-komiser-600 checked:hover:border-transparent checked:hover:bg-komiser-700"
    />
  );
}

export default Checkbox;
