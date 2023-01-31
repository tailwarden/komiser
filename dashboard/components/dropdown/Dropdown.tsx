import { ReactNode } from 'react';

type DropdownProps = {
  isOpen: boolean;
  toggle: () => void;
  children: ReactNode;
};

function Dropdown({ isOpen, toggle, children }: DropdownProps) {
  return (
    <button
      className={`flex h-[2.5rem] items-center gap-2 rounded-lg border-2 border-primary px-3 text-sm font-medium text-primary transition-colors active:border-primary active:text-primary ${
        isOpen
          ? 'bg-komiser-200/50'
          : 'bg-transparent hover:bg-komiser-200/30 active:bg-komiser-200'
      }`}
      onClick={toggle}
    >
      {children}
    </button>
  );
}

export default Dropdown;
