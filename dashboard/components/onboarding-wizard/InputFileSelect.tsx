import { MutableRefObject, ReactNode } from 'react';

interface InputFileSelectProps {
  id: string;
  value?: any;
  type: string;
  label: string;
  icon?: ReactNode;
  subLabel?: string;
  disabled?: boolean;
  placeholder?: string;
  iconClick?: () => void;
  handleFileChange: (event: any) => void;
  fileInputRef: MutableRefObject<HTMLInputElement | null>;
}

function InputFileSelect({
  id,
  icon,
  type,
  label,
  value,
  subLabel,
  iconClick,
  placeholder,
  fileInputRef,
  handleFileChange,
  disabled = false
}: InputFileSelectProps) {
  return (
    <div>
      <label htmlFor={id} className="mb-2 block text-gray-700">
        {label}
      </label>

      {subLabel && (
        <span className="-mt-[5px] mb-2 block text-xs leading-4 text-black-400">
          {subLabel}
        </span>
      )}

      <div className="relative mb-6">
        <input
          type="file"
          className="hidden"
          ref={fileInputRef}
          onChange={handleFileChange}
        />

        <input
          id={id}
          type={type}
          value={value}
          disabled={disabled}
          placeholder={placeholder}
          className="block w-full rounded border py-4 pl-5 text-sm text-black-900 outline outline-black-200 focus:outline-2 focus:outline-primary"
        />

        {icon && (
          <button
            onClick={iconClick}
            className="absolute inset-y-0 right-5 flex items-center pl-3 text-komiser-600"
          >
            {icon}
          </button>
        )}
      </div>
    </div>
  );
}

export default InputFileSelect;
