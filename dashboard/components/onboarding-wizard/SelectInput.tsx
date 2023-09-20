import { ReactNode, useState } from 'react';
import classNames from 'classnames';

import ChevronDownIcon from '../icons/ChevronDownIcon';

export type SelectInputProps = {
  label: string;
  value: string;
  values: string[];
  icon?: string | ReactNode;
  handleChange: (value: string) => void;
  displayValues: { [key: string]: any }[];
};

function SelectInput({
  label,
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
        right-4 text-komiser-600 transition-all"
      >
        {icon}
      </div>

      <label className="mb-2 block text-gray-700">{label}</label>
      <button
        onClick={toggle}
        className={classNames(
          'h-[60px] w-full overflow-hidden rounded bg-white text-left outline outline-black-200/50 hover:outline-black-200 focus:outline-primary',
          { 'outline-2 outline-primary': isOpen }
        )}
      >
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pb-[0.75rem] pl-4 pr-16 pt-[0.75rem] text-sm text-black-900">
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
          <div className="absolute top-[96px] z-[21] max-h-52 w-full overflow-hidden overflow-y-auto rounded-lg border border-black-130 bg-white px-3 py-2 shadow-lg">
            <div className="flex w-full flex-col gap-1">
              {values.map((item, idx) => {
                const isActive = value === item;
                return (
                  <button
                    key={idx}
                    className={classNames(
                      'flex items-center rounded px-3 py-2 text-left text-sm text-black-400 hover:bg-black-150',
                      { 'bg-komiser-150': isActive }
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
