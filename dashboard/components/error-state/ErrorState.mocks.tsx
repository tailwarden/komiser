import Button from '../button/Button';
import type { ErrorStateProps } from './ErrorState';

const base: ErrorStateProps = {
  title: 'Network request error',
  message:
    'There was an error fetching the cloud accounts. Please refer to the logs for more info and try again.',
  action: (
    <Button style="outline" size="lg">
      Refresh the page
    </Button>
  )
};

const mockErrorStateProps = {
  base
};

export default mockErrorStateProps;
