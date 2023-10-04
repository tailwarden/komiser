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
  betaFlagOnboardingWizard: boolean;
};

export const initialContext = {
  displayBanner: false,
  dismissBanner: () => {},
  loading: false,
  data: undefined,
  error: false,
  hasNoAccounts: false,
  fetch: () => {},
  betaFlagOnboardingWizard: false
};

const GlobalAppContext = createContext<GlobalAppContextProps>(initialContext);

export default GlobalAppContext;
