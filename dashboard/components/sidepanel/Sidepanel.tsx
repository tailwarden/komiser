import { ReactNode } from 'react';

type SidepanelProps = {
  isOpen: boolean;
  closeModal: () => void;
  children: ReactNode;
};

function Sidepanel({ isOpen, closeModal, children }: SidepanelProps) {
  return (
    <>
      {isOpen && (
        <>
          <div
            onClick={closeModal}
            className="hidden sm:block fixed inset-0 z-30 bg-black-900/10 opacity-0 animate-fade-in"
          ></div>
          <div className="fixed overflow-auto inset-0 z-30 sm:top-4 sm:bottom-4 sm:right-4 sm:left-auto w-full sm:w-[32rem] p-6 sm:rounded-lg shadow-2xl opacity-0 animate-fade-in-up sm:animate-fade-in-left bg-white">
            {children}
          </div>
        </>
      )}
    </>
  );
}

export default Sidepanel;
