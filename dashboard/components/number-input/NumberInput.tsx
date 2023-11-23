import { ChangeEvent, KeyboardEvent, useEffect, useRef, useState } from 'react';
import MinusIcon from '@components/icons/MinusIcon';
import PlusIcon from '@components/icons/PlusIcon';
import { required } from '../../utils/regex';

export type InputEvent = ChangeEvent<HTMLInputElement>;

export type InputProps = {
  disabled?: boolean;
  id?: number;
  name: string;
  label?: string;
  required?: boolean;
  regex?: RegExp;
  error?: string;
  value: number;
  autofocus?: boolean;
  min?: number;
  max?: number;
  maxLength?: number;
  positiveNumberOnly?: boolean;
  action: (newData: any, id?: number) => void;
  handleValueChange: (value: number) => void;
  step?: number;
};

function NumberInput({
  id,
  name,
  label,
  regex = required,
  error = 'Please provide a value',
  autofocus,
  positiveNumberOnly,
  action,
  handleValueChange,
  value,
  step = 1,
  maxLength,
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

  const adjustBtn = `absolute ${
    label ? 'w-14' : 'w-11'
  } h-full p-3 border-gray-200 inline-flex justify-center items-center focus:outline-none`;

  const iconStyle = `text-neutral-900 ${label ? 'w-8 h-8' : 'w-6 h-6'}`;

  return (
    <div>
      <div className={`relative flex w-full ${label ? 'h-14' : 'h-11'}`}>
        <button
          className={`${adjustBtn} left-0 border-r`}
          onClick={() => handleValueChange(value - step)}
        >
          <MinusIcon className={iconStyle} />
        </button>
        <input
          name={name}
          className={`peer w-full rounded bg-white px-12 py-[0.75rem] ${
            label && 'pt-[1.75rem]'
          } text-neutral-900 text-center text-sm caret-darkcyan-500 outline outline-[0.063rem] outline-gray-300 focus:outline-[0.12rem] focus:outline-darkcyan-500 ${
            isValid === false && `outline-red-500 focus:outline-red-500`
          }`}
          type="number"
          placeholder=" "
          onFocus={handleFocus}
          onBlur={e => handleBlur(e)}
          onChange={e => {
            // e.target.value = e.target.value.slice(0, maxLength)
            // if(Number(e.target.value) === 0) e.target.value = "0"
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
          value={value}
          step={step}
          {...otherProps}
        />
        <button
          className={`${adjustBtn} right-0 border-l`}
          onClick={() => handleValueChange(value + step)}
        >
          <PlusIcon className={iconStyle} />
        </button>
        {label && (
          <span className="font-['Noto Sans'] text-neutral-500 pointer-events-none absolute bottom-[1.925rem] left-1/2 -translate-x-1/2 select-none text-xs font-normal transition-all peer-placeholder-shown:bottom-[1.15rem] peer-placeholder-shown:left-4 peer-placeholder-shown:scale-[87.5%] peer-focus:bottom-[1.925rem]">
            {label}
          </span>
        )}
      </div>
      {isValid === false && (
        <p className="mt-2 text-xs text-red-500">{error}</p>
      )}
    </div>
  );
}

export default NumberInput;
