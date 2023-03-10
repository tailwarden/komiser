import { createContext } from 'react';

export type GlobalData = {
  resources: number;
  regions: number;
  costs: number;
  accounts: number;
};

export type GlobalAppContextProps = {
  displayBanner: boolean;
  dismissBanner: () => void;
  loading: boolean;
  data: GlobalData | undefined;
  error: boolean;
  hasNoAccounts: boolean;
  fetch: () => void;
};

const GlobalAppContext = createContext<GlobalAppContextProps>({
  displayBanner: false,
  dismissBanner: () => {},
  loading: true,
  data: undefined,
  error: false,
  hasNoAccounts: false,
  fetch: () => {}
});

export default GlobalAppContext;
