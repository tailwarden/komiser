import { render, screen } from '@testing-library/react';
import DashboardTopStats from './DashboardTopStats';
import GlobalAppContext from '../../../layout/context/GlobalAppContext';

const initialContext = {
  displayBanner: false,
  dismissBanner: () => {},
  loading: false,
  data: undefined,
  error: false,
  hasNoAccounts: false,
  fetch: () => {}
};

describe('Dashboard Top Stats', () => {
  test('should render the skeleton component when loading is true', () => {
    render(
      <GlobalAppContext.Provider
        value={{
          ...initialContext,
          loading: true
        }}
      >
        <DashboardTopStats />
      </GlobalAppContext.Provider>
    );
    const skeleton = screen.getByTestId('loading');
    expect(skeleton).toBeInTheDocument();
  });

  test('should render the error component when error is true', () => {
    render(
      <GlobalAppContext.Provider
        value={{
          ...initialContext,
          error: true
        }}
      >
        <DashboardTopStats />
      </GlobalAppContext.Provider>
    );
    const error = screen.getByTestId('error');
    expect(error).toBeInTheDocument();
  });

  test('should render the component if error and loading are false', () => {
    render(
      <GlobalAppContext.Provider
        value={{
          ...initialContext,
          data: { resources: 25, regions: 17, costs: 5, accounts: 20 }
        }}
      >
        <DashboardTopStats />
      </GlobalAppContext.Provider>
    );
    const component = screen.getByTestId('data');
    expect(component).toBeInTheDocument();
  });
});
