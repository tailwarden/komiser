import { useState } from 'react';

function useViews() {
  const [isOpen, setIsOpen] = useState(false);

  return {
    isOpen
  };
}

export default useViews;
