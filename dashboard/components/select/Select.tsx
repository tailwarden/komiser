import { ChangeEvent } from 'react';

export type SelectProps = {
  label: string;
  value?: string | number;
  displayValues?: string[];
  onChange: (e: ChangeEvent<HTMLSelectElement>) => void;
  values: string[] | number[];
  chevronPadding?: boolean;
};

function Select({
  label,
  value,
  displayValues,
  onChange,
  values,
  chevronPadding
}: SelectProps) {
  return (
    <div className="relative">
      <div
        className={`pointer-events-none absolute bottom-[1.15rem] ${
          chevronPadding ? 'right-4' : 'right-2'
        }  text-black-200 transition-all`}
      >
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M7 10L12 15L17 10"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </div>
      <span className="pointer-events-none absolute left-4 bottom-[1.925rem] origin-left scale-75 select-none font-normal text-black-300">
        {label}
      </span>
      <select
        value={value}
        onChange={e => onChange(e)}
        className="w-full appearance-none rounded bg-white pt-[1.75rem] pb-[0.75rem] pl-4 pr-16 text-sm text-black-900 caret-primary outline outline-black-200 hover:outline-black-300 focus:outline-2 focus:outline-primary"
      >
        {values.map((item, idx) => (
          <option key={idx} value={item}>
            {displayValues ? displayValues[idx] : item}
          </option>
        ))}
      </select>
    </div>
  );
}

export default Select;
