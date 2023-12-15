import { ReactNode, ChangeEvent } from 'react';

interface LabelledInputProps {
  id: string;
  name?: string;
  value?: any;
  type: string;
  label: string;
  icon?: ReactNode;
  subLabel?: string;
  disabled?: boolean;
  placeholder?: string;
  required?: boolean;
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void;
}

function LabelledInput({
  id,
  name,
  icon,
  type,
  label,
  value,
  subLabel,
  placeholder,
  required = false,
  disabled = false,
  onChange
}: LabelledInputProps) {
  return (
    <div>
      <label htmlFor={id} className="mb-2 block text-gray-700">
        {label}
      </label>

      {subLabel && (
        <span className="-mt-[5px] mb-2 block text-xs leading-4 text-gray-700">
          {subLabel}
        </span>
      )}

      <div className="relative mb-6">
        {icon && (
          <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-gray-500">
            {icon}
          </div>
        )}

        <input
          id={id}
          type={type}
          name={name ?? id}
          disabled={disabled}
          placeholder={placeholder}
          required={required}
          className={`block w-full rounded py-[14.5px] text-sm text-gray-950 outline outline-gray-300 focus:outline-2 focus:outline-darkcyan-500 ${
            icon ? 'pl-10' : 'pl-3'
          }`}
          onChange={onChange}
          defaultValue={value}
          autoComplete="off"
        />
      </div>
    </div>
  );
}

export default LabelledInput;
