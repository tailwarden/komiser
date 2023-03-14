import { render, screen } from '@testing-library/react';
import ExportCSVButton from './ExportCSVButton';

const props = {
  id: undefined,
  exportCSV: jest.fn(),
  loading: false
};

describe('ExportCSV component', () => {
  it('should render without crashing', () => {
    render(<ExportCSVButton {...props} />);
  });
});
