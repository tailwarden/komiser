import { createContext } from 'react';

type LayoutContextProps = {
  displayBanner: boolean;
  dismissBanner: () => void;
};

const LayoutContext = createContext<LayoutContextProps>({
  displayBanner: false,
  dismissBanner: () => {}
});

export default LayoutContext;
