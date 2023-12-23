import { ReactNode, useEffect } from 'react';

export type ModalProps = {
  isOpen: boolean;
  closeModal: () => void;
  children: ReactNode;
  id?: string;
};

function Modal({ isOpen, closeModal, id, children }: ModalProps) {
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
            id={`${id}-wrapper`}
            onClick={closeModal}
            className="fixed inset-0 z-30 hidden animate-fade-in bg-gray-900/60 sm:block"
          ></div>
          <div
            id={`${id}-modal`}
            className="fixed inset-0 z-30 w-full animate-fade-in-down-short bg-white p-8 opacity-0 shadow-left sm:bottom-auto sm:top-[15%] sm:m-auto sm:max-w-fit sm:rounded-lg"
          >
            {children}
          </div>
        </>
      )}
    </>
  );
}

export default Modal;
