import { ReactNode, useEffect } from 'react';

type ModalProps = {
  isOpen: boolean;
  closeModal: () => void;
  children: ReactNode;
};

function Modal({ isOpen, closeModal, children }: ModalProps) {
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
            className="fixed inset-0 z-30 hidden animate-fade-in bg-black-900/10 opacity-0 sm:block"
          ></div>
          <div className="fixed inset-0 z-30 w-full animate-fade-in-down-short bg-white p-8 opacity-0 shadow-2xl sm:top-[15%] sm:bottom-auto sm:m-auto sm:max-w-fit sm:rounded-lg">
            {children}
          </div>
        </>
      )}
    </>
  );
}

export default Modal;
