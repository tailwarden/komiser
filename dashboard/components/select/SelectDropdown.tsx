import { useState } from 'react';

type SelectDropdownProps = {
  label: string;
  value: string;
  values: string[];
  displayValues: string[];
  handleChange: (value: string) => void;
};

function SelectDropdown({
  label,
  value,
  values,
  displayValues,
  handleChange
}: SelectDropdownProps) {
  const [isOpen, setIsOpen] = useState(false);
  const index = values.findIndex(currentValue => currentValue === value);

  function toggle() {
    setIsOpen(!isOpen);
  }

  return (
    <div className="relative">
      <div
        className={`pointer-events-none absolute right-2
        bottom-[1.15rem] text-black-200 transition-all`}
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
      <button
        onClick={toggle}
        className={`h-[60px] w-full overflow-hidden rounded text-left outline hover:outline-black-200 focus:outline-2 focus:outline-primary ${
          isOpen ? 'outline-2 outline-primary' : 'outline-black-200/50'
        }`}
      >
        <div className="absolute right-0 top-1 h-[50px] w-6"></div>
        <span className="pointer-events-none absolute left-4 bottom-[1.925rem] origin-left scale-75 select-none font-normal text-black-300">
          {label}
        </span>
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pt-[1.75rem] pb-[0.75rem] pl-4 pr-16 text-sm text-black-900">
          {displayValues[index]}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[4.15rem] z-[21] w-full overflow-hidden rounded-lg border border-black-200 bg-white p-2 shadow-lg">
            <div className="flex w-full flex-col">
              {values.map((item, idx) => {
                const isActive = value === item;
                return (
                  <button
                    key={idx}
                    className={`flex items-center justify-between rounded-lg p-3 text-left text-sm hover:bg-black-100 ${
                      isActive ? 'text-primary' : 'text-black-400'
                    }`}
                    onClick={() => handleChange(item)}
                  >
                    {displayValues[idx]}
                    {isActive && (
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        fill="none"
                        viewBox="0 0 24 24"
                      >
                        <path
                          stroke="currentColor"
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="1.5"
                          d="M7 12l3.662 4L18 8"
                        ></path>
                      </svg>
                    )}
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

export default SelectDropdown;
