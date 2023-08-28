import { ChangeEvent, KeyboardEvent, useEffect, useRef, useState } from 'react';
import { required } from '../../utils/regex';

export type InputEvent = ChangeEvent<HTMLInputElement>;

export type InputProps = {
  id?: number;
  name: string;
  type?: string;
  label: string;
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
  type = 'text',
  label,
  regex = required,
  error = 'Please provide a value',
  value,
  autofocus,
  min,
  maxLength,
  positiveNumberOnly,
  action
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
          type={type}
          name={name}
          className={`peer w-full rounded bg-white px-4 pb-[0.75rem] pt-[1.75rem] text-sm text-black-900 caret-primary outline outline-[0.063rem] outline-black-200 focus:outline-[0.12rem] focus:outline-primary ${
            isValid === false && `outline-error-600 focus:outline-error-600`
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
          value={value}
          ref={inputRef}
          min={min}
          maxLength={maxLength}
          autoComplete="off"
          data-lpignore="true"
          data-form-type="other"
        />
        <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-black-300 transition-all peer-placeholder-shown:bottom-[1.15rem] peer-placeholder-shown:left-4 peer-placeholder-shown:scale-[87.5%] peer-focus:bottom-[1.925rem] peer-focus:scale-75">
          {label}
        </span>
      </div>
      {isValid === false && (
        <p className="mt-2 text-xs text-error-600">{error}</p>
      )}
    </div>
  );
}

export default Input;
