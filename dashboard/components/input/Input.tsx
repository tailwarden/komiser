import { ChangeEvent, useState } from 'react';

export type InputEvent = ChangeEvent<HTMLInputElement>;

export type InputProps = {
  id?: number;
  name: string;
  type: string;
  label: string;
  regex: RegExp;
  error: string;
  value?: string | number | string[];
  action: (newData: any, id?: number) => void;
};

function Input({
  id,
  name,
  type,
  label,
  regex,
  error,
  value,
  action
}: InputProps) {
  const [isValid, setIsValid] = useState<boolean | undefined>(undefined);

  function handleBlur(e: InputEvent): void {
    const trimmedValue = e.target.value.trim();
    if (!regex || !trimmedValue) return;

    const testResult = regex.test(trimmedValue);
    setIsValid(testResult);
  }

  function handleFocus(): void {
    setIsValid(undefined);
  }

  return (
    <div>
      <div className="relative">
        <input
          type={type}
          name={name}
          className={`w-full pt-[1.75rem] pb-[0.75rem] px-4 text-sm bg-white text-black-900 rounded caret-secondary outline outline-black-200 focus:outline-secondary focus:outline-2 peer ${
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
          value={value}
          autoComplete="off"
          data-lpignore="true"
          data-form-type="other"
        />
        <span className="absolute select-none scale-75 left-4 bottom-[1.925rem] pointer-events-none text-black-300 font-normal peer-focus:scale-75 peer-focus:bottom-[1.925rem] peer-placeholder-shown:scale-[87.5%] peer-placeholder-shown:left-4 peer-placeholder-shown:bottom-[1.15rem] origin-left transition-all">
          {label}
        </span>
      </div>
      {isValid === false && (
        <p className="mt-2 text-error-600 text-xs">{error}</p>
      )}
    </div>
  );
}

export default Input;
