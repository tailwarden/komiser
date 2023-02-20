import { useEffect, useState } from 'react';

function useCloudMapExpand() {
  const [isOpen, setIsOpen] = useState(false);

  function toggle() {
    setIsOpen(!isOpen);
  }

  function close() {
    setIsOpen(false);
  }

  useEffect(() => {
    function escFunction(event: KeyboardEvent) {
      if (event.key === 'Escape') {
        close();
      }
    }

    document.addEventListener('keydown', escFunction, false);

    return () => {
      document.removeEventListener('keydown', escFunction, false);
    };
  }, []);

  return { isOpen, toggle };
}

export default useCloudMapExpand;
