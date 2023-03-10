import { render, screen } from '@testing-library/react';
import { ReactElement } from 'react';
import GlobalAppContext, {
  initialContext,
  GlobalAppContextProps
} from '../../../layout/context/GlobalAppContext';
import mockDataForDashboard from '../../utils/mockDataForDashboard';
import DashboardTopStats from './DashboardTopStats';

const customRender = (
  children: ReactElement,
  props?: Partial<GlobalAppContextProps>
) =>
  render(
    <GlobalAppContext.Provider value={{ ...initialContext, ...props }}>
      {children}
    </GlobalAppContext.Provider>
  );

describe('Dashboard Top Stats', () => {
  it('should render without crashing', () => {
    customRender(<DashboardTopStats />);
  });

  it('should render the skeleton component when loading is true', () => {
    customRender(<DashboardTopStats />, { loading: true });
    const skeleton = screen.getByTestId('loading');
    expect(skeleton).toBeInTheDocument();
  });

  it('should render the error component when error is true', () => {
    customRender(<DashboardTopStats />, { error: true });
    const error = screen.getByTestId('error');
    expect(error).toBeInTheDocument();
  });

  it('should render the top stats cards component when there is data', () => {
    customRender(<DashboardTopStats />, { data: mockDataForDashboard.stats });
    const component = screen.getByTestId('data');
    expect(component).toBeInTheDocument();
  });
});
