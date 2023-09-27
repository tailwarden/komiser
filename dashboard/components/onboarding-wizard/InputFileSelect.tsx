import { MutableRefObject, ReactNode } from 'react';
import classNames from 'classnames';
import AlertCircleIcon from '../icons/AlertCircleIcon';
import AlertIcon from '../icons/AlertIcon';
import AlertCircleIconFilled from '../icons/AlertCircleIconFilled';

interface InputFileSelectProps {
  id: string;
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
        <span className="-mt-[5px] mb-2 block text-xs leading-4 text-black-400">
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
          type={type}
          value={value}
          disabled={disabled}
          placeholder={placeholder}
          className={classNames(
            hasError
              ? 'outline-error-600 focus:outline-error-700'
              : 'outline-gray-200 focus:outline-primary',
            'block w-full rounded border py-4 pl-5 text-sm text-black-900 outline focus:outline-2 '
          )}
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
      {hasError && errorMessage && (
        <div className="mt-2 flex items-center text-sm text-error-600">
          <AlertCircleIconFilled className="mr-1 inline-block h-4 w-4" />
          {errorMessage}
        </div>
      )}
    </div>
  );
}

export default InputFileSelect;
