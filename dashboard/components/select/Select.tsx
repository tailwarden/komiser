import { useState } from 'react';
import classNames from 'classnames';
import ChevronDownIcon from '../icons/ChevronDownIcon';

export type SelectProps = {
  label: string;
  value: string;
  values: string[];
  displayValues: string[];
  handleChange: (value: string) => void;
};

function Select({
  label,
  value,
  values,
  displayValues,
  handleChange
}: SelectProps) {
  const [isOpen, setIsOpen] = useState(false);
  const index = values.findIndex(currentValue => currentValue === value);

  function toggle() {
    setIsOpen(!isOpen);
  }

  return (
    <div className="relative">
      <div
        className="pointer-events-none absolute bottom-[1.15rem]
        right-4 text-black-900 transition-all"
      >
        <ChevronDownIcon width={24} height={24} />
      </div>
      <button
        onClick={toggle}
        className={classNames(
          'h-[60px] w-full overflow-hidden rounded text-left outline outline-black-200/50 hover:outline-black-200 focus:outline-primary',
          { 'outline-2 outline-primary': isOpen }
        )}
      >
        <div className="absolute right-0 top-1 h-[50px] w-6"></div>
        <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-black-300">
          {label}
        </span>
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pb-[0.75rem] pl-4 pr-16 pt-[1.75rem] text-sm text-black-900">
          {displayValues[index]}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[66px] z-[21] max-h-52 w-full overflow-hidden overflow-y-auto rounded-lg border border-black-130 bg-white py-2 px-3 shadow-lg">
            <div className="flex w-full flex-col gap-1">
              {values.map((item, idx) => {
                const isActive = value === item;
                return (
                  <button
                    key={idx}
                    className={classNames(
                      'flex items-center justify-between rounded px-3 py-2 text-left text-sm text-black-400 hover:bg-black-150',
                      { 'bg-komiser-150': isActive }
                    )}
                    onClick={() => handleChange(item)}
                  >
                    {displayValues[idx]}
                  </button>
                );
              })}
            </div>
          </div>
        </>
      )}
    </div>
  );
}

export default Select;
