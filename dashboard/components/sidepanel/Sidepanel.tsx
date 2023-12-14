import { ReactNode, useEffect } from 'react';

type SidepanelProps = {
  isOpen: boolean;
  closeModal: () => void;
  children: ReactNode;
  noScroll?: boolean;
};

function Sidepanel({ isOpen, closeModal, children, noScroll }: SidepanelProps) {
  // Listen to ESC key and close modal
  useEffect(() => {
    function escFunction(event: KeyboardEvent) {
      if (event.key === 'Escape') {
        closeModal();
      }
    }

    document.addEventListener('keydown', escFunction, false);

    return () => {
      document.removeEventListener('keydown', escFunction, false);
    };
  }, []);

  return (
    <>
      {isOpen && (
        <>
          <div
            onClick={closeModal}
            className="fixed inset-0 z-30 hidden animate-fade-in bg-gray-900/30  sm:block"
          ></div>
          <div
            className={`fixed inset-0 z-30 flex w-full animate-fade-in-up flex-col gap-4 overflow-auto bg-white px-8 pt-4 opacity-0 shadow-left sm:bottom-4 sm:left-auto sm:right-4 sm:top-4 sm:w-[38rem] sm:animate-fade-in-left sm:rounded-lg ${
              noScroll ? 'overflow-hidden' : 'overflow-auto'
            }`}
          >
            {children}
          </div>
        </>
      )}
    </>
  );
}

export default Sidepanel;
