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
        right-4 text-gray-950 transition-all"
      >
        <ChevronDownIcon width={24} height={24} />
      </div>
      <button
        onClick={toggle}
        className={classNames(
          'h-[60px] w-full overflow-hidden rounded text-left outline outline-gray-300 hover:outline-gray-300 focus:outline-darkcyan-500',
          { 'outline-2 outline-darkcyan-500': isOpen }
        )}
      >
        <div className="absolute right-0 top-1 h-[50px] w-6"></div>
        <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-gray-500">
          {label}
        </span>
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pb-[0.75rem] pl-4 pr-16 pt-[1.75rem] text-sm text-gray-950">
          {displayValues[index]}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[66px] z-[21] max-h-52 w-full overflow-hidden overflow-y-auto rounded-lg border border-gray-100 bg-white px-3 py-2 shadow-right">
            <div className="flex w-full flex-col gap-1">
              {values.map((item, idx) => {
                const isActive = value === item;
                return (
                  <button
                    key={idx}
                    className={classNames(
                      'flex items-center justify-between rounded px-3 py-2 text-left text-sm text-gray-700 hover:bg-background-ds',
                      { 'bg-cyan-100': isActive }
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
