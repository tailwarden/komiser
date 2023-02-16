import { createContext } from 'react';

type GlobalAppContextProps = {
  displayBanner: boolean;
  dismissBanner: () => void;
  loading: boolean;
  data: any;
  error: boolean;
  hasNoAccounts: boolean;
  fetch: () => void;
};

const GlobalAppContext = createContext<GlobalAppContextProps>({
  displayBanner: false,
  dismissBanner: () => {},
  loading: true,
  data: {},
  error: false,
  hasNoAccounts: false,
  fetch: () => {}
});

export default GlobalAppContext;
