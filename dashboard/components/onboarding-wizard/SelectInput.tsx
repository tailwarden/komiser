import { ReactNode, useState } from 'react';
import classNames from 'classnames';

import ChevronDownIcon from '../icons/ChevronDownIcon';

export type SelectInputProps = {
  label: string;
  name?: string;
  value: string;
  values: string[];
  icon?: string | ReactNode;
  handleChange: (value: string) => void;
  displayValues: { [key: string]: any }[];
};

function SelectInput({
  label,
  name,
  value,
  values,
  handleChange,
  displayValues,
  icon = <ChevronDownIcon width={24} height={24} />
}: SelectInputProps) {
  const [isOpen, setIsOpen] = useState(false);
  const index = values.findIndex(currentValue => currentValue === value);

  function toggle() {
    setIsOpen(!isOpen);
  }

  function handleClick(item: string) {
    handleChange(item);
    toggle();
  }

  return (
    <div className="relative">
      <div
        className="pointer-events-none absolute bottom-[1.15rem]
        right-4 text-darkcyan-500 transition-all"
      >
        {icon}
      </div>

      <input type="hidden" name={name} value={value} readOnly />

      <label className="mb-2 block text-gray-700">{label}</label>
      <button
        onClick={toggle}
        className={classNames(
          'h-[60px] w-full overflow-hidden rounded bg-white text-left outline outline-gray-300 hover:outline-gray-300 focus:outline-darkcyan-500',
          { 'outline-2 outline-darkcyan-500': isOpen }
        )}
      >
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pb-[0.75rem] pl-4 pr-16 pt-[0.75rem] text-sm text-gray-950">
          {displayValues[index].icon && displayValues[index].icon}
          {displayValues[index].label}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[96px] z-[21] max-h-52 w-full overflow-hidden overflow-y-auto rounded-lg border border-gray-100 bg-white px-3 py-2 shadow-right">
            <div className="flex w-full flex-col gap-1">
              {values.map((item, idx) => {
                const isActive = value === item;
                return (
                  <button
                    key={idx}
                    className={classNames(
                      'flex items-center rounded px-3 py-2 text-left text-sm text-gray-700 hover:bg-background-ds',
                      { 'bg-cyan-100': isActive }
                    )}
                    onClick={() => handleClick(item)}
                  >
                    {displayValues[idx].icon && displayValues[idx].icon}
                    <div className={displayValues[idx].icon ? 'pl-3' : ''}>
                      {displayValues[idx].label}
                    </div>
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

export default SelectInput;
