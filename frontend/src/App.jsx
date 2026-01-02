import { useState } from 'react';
import Container from './components/Container';
import FilterTabs from './components/FilterTabs';
import CountdownDisplay from './components/CountdownDisplay';
import ProgressBar from './components/ProgressBar';
import { useCountdown } from './hooks/useCountdown';

function App() {
  const [filter, setFilter] = useState('year');
  const { data, loading, error } = useCountdown(filter);

  return (
    <div className="app">
      <Container>
        <div className="arcade-header">
          <h1 className="arcade-title">COUNTDOWN</h1>
        </div>

        <FilterTabs activeFilter={filter} onFilterChange={setFilter} />

        {loading && <div className="loading">LOADING...</div>}

        {error && <div className="error">ERROR: {error}</div>}

        {!loading && !error && data && (
          <>
            <CountdownDisplay data={data} filter={filter} />
            <ProgressBar percentComplete={data.percentComplete} />
          </>
        )}
      </Container>
    </div>
  );
}

export default App;
