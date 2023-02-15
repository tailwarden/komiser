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
          <div className="fixed inset-0 z-30 w-full animate-fade-in-down-short overflow-auto bg-white p-6 opacity-0 shadow-2xl sm:top-[30%] sm:bottom-[70%] sm:m-auto sm:min-h-fit sm:w-[28rem] sm:rounded-lg">
            {children}
          </div>
        </>
      )}
    </>
  );
}

export default Modal;
