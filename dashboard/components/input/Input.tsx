import { ChangeEvent, KeyboardEvent, useEffect, useRef, useState } from 'react';
import { required } from '../../utils/regex';

export type InputEvent = ChangeEvent<HTMLInputElement>;

export type InputProps = {
  disabled?: boolean;
  id?: number;
  name: string;
  type: string;
  label: string;
  required?: boolean;
  regex?: RegExp;
  error?: string;
  value?: string | number | string[];
  autofocus?: boolean;
  min?: number;
  maxLength?: number;
  positiveNumberOnly?: boolean;
  action: (newData: any, id?: number) => void;
};

function Input({
  id,
  name,
  label,
  regex = required,
  error = 'Please provide a value',
  autofocus,
  positiveNumberOnly,
  action,
  ...otherProps
}: InputProps) {
  const [isValid, setIsValid] = useState<boolean | undefined>(undefined);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (autofocus) {
      inputRef.current?.focus();
    }
  }, []);

  function handleBlur(e: InputEvent): void {
    const trimmedValue = e.target.value.trim();
    if (!regex || !trimmedValue) return;

    const testResult = regex.test(trimmedValue);
    setIsValid(testResult);
  }

  function handleFocus(): void {
    setIsValid(undefined);
  }

  function handleKeyDown(e: KeyboardEvent) {
    if (positiveNumberOnly) {
      const invalidChars = ['-', '+', 'e'];
      if (invalidChars.includes(e.key)) {
        e.preventDefault();
      }
    }
  }

  return (
    <div>
      <div className="relative">
        <input
          name={name}
          className={`peer w-full rounded bg-white px-4 pb-[0.75rem] pt-[1.75rem] text-sm text-gray-950 caret-darkcyan-500 outline outline-[0.063rem] outline-gray-300 focus:outline-[0.12rem] focus:outline-darkcyan-500 ${
            isValid === false && `outline-red-500 focus:outline-red-500`
          }`}
          placeholder=" "
          onFocus={handleFocus}
          onBlur={e => handleBlur(e)}
          onChange={e => {
            if (typeof id === 'number') {
              action({ [name]: e.target.value }, id);
            } else {
              action({ [name]: e.target.value });
            }
          }}
          onKeyDown={e => handleKeyDown(e)}
          ref={inputRef}
          autoComplete="off"
          data-lpignore="true"
          data-form-type="other"
          {...otherProps}
        />
        <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-gray-500 transition-all peer-placeholder-shown:bottom-[1.15rem] peer-placeholder-shown:left-4 peer-placeholder-shown:scale-[87.5%] peer-focus:bottom-[1.925rem] peer-focus:scale-75">
          {label}
        </span>
      </div>
      {isValid === false && (
        <p className="mt-2 text-xs text-red-500">{error}</p>
      )}
    </div>
  );
}

export default Input;
