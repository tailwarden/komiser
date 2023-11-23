import { MutableRefObject, ReactNode } from 'react';
import classNames from 'classnames';
import AlertCircleIcon from '../icons/AlertCircleIcon';
import AlertIcon from '../icons/AlertIcon';
import AlertCircleIconFilled from '../icons/AlertCircleIconFilled';

interface InputFileSelectProps {
  id: string;
  name?: string;
  value?: any;
  type: string;
  label: string;
  icon?: ReactNode;
  subLabel?: string;
  disabled?: boolean;
  hasError?: boolean;
  errorMessage?: string;
  placeholder?: string;
  iconClick?: () => void;
  handleFileChange: (event: any) => void;
  handleInputChange: (event: any) => void;
  fileInputRef: MutableRefObject<HTMLInputElement | null>;
}

function InputFileSelect({
  id,
  name,
  icon,
  type,
  label,
  value,
  subLabel,
  iconClick,
  placeholder,
  fileInputRef,
  handleFileChange,
  handleInputChange,
  disabled = false,
  hasError = false,
  errorMessage = ''
}: InputFileSelectProps) {
  return (
    <div className="relative mb-6">
      <label htmlFor={id} className="mb-2 block text-gray-700">
        {label}
      </label>

      {subLabel && (
        <span className="-mt-[5px] mb-2 block text-xs leading-4 text-gray-700">
          {subLabel}
        </span>
      )}

      <div className="relative">
        <input
          type="file"
          className="hidden"
          ref={fileInputRef}
          onChange={handleFileChange}
        />

        <input
          id={id}
          name={name ?? id}
          type={type}
          value={value}
          disabled={disabled}
          placeholder={placeholder}
          className={classNames(
            hasError
              ? 'outline-red-500 focus:outline-red-500'
              : 'outline-gray-200 focus:outline-darkcyan-500',
            'block w-full rounded border py-4 pl-5 text-sm text-gray-950 outline focus:outline-2 '
          )}
          onChange={handleInputChange}
        />

        {icon && (
          <button
            onClick={iconClick}
            className="absolute inset-y-0 right-5 flex items-center pl-3 text-darkcyan-500"
          >
            {icon}
          </button>
        )}
      </div>
      {hasError && errorMessage && (
        <div className="mt-2 flex items-center text-sm text-red-500">
          <AlertCircleIconFilled className="mr-1 inline-block h-4 w-4" />
          {errorMessage}
        </div>
      )}
    </div>
  );
}

export default InputFileSelect;
