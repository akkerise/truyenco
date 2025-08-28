import React, { useCallback, useEffect, useMemo, useState, Suspense } from 'react';
import { AppProvider, useApp } from './context/AppContext';
import { useApi } from './hooks/useApi';
import { formatDate } from './utils/date';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';

const LazyModal = React.lazy(() => import('./components/Modal'));

const Counter = React.memo(() => {
  const { state, dispatch } = useApp();
  const inc = useCallback(() => dispatch({ type: 'increment' }), [dispatch]);
  const dec = useCallback(() => dispatch({ type: 'decrement' }), [dispatch]);
  const doubled = useMemo(() => state.count * 2, [state.count]);
  return (
    <div className="space-x-2">
      <span>Count: {state.count}</span>
      <span>Doubled: {doubled}</span>
      <button onClick={inc} className="btn">+</button>
      <button onClick={dec} className="btn">-</button>
    </div>
  );
});

const schema = yup.object({
  email: yup.string().required().email(),
});

type FormValues = { email: string };

const Form = React.memo(() => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormValues>({ resolver: yupResolver(schema) });
  const onSubmit = useCallback((data: FormValues) => {
    console.log(data);
  }, []);
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-2">
      <input {...register('email')} className="border" placeholder="email" />
      {errors.email && <span className="text-red-500">{errors.email.message}</span>}
      <button type="submit" className="btn">Submit</button>
    </form>
  );
});

function AppContent() {
  const api = useApi();
  useEffect(() => { api.get('/health').catch(() => {}); }, [api]);
  const today = useMemo(() => formatDate(new Date()), []);
  const [open, setOpen] = useState(false);
  const openModal = useCallback(() => setOpen(true), []);
  const closeModal = useCallback(() => setOpen(false), []);
  return (
    <div className="p-4 space-y-4">
      <div>Today: {today}</div>
      <Counter />
      <Form />
      <button onClick={openModal} className="btn">Open Modal</button>
      {open && (
        <Suspense fallback={null}>
          <LazyModal onClose={closeModal} />
        </Suspense>
      )}
    </div>
  );
}

function App() {
  return (
    <AppProvider>
      <AppContent />
    </AppProvider>
  );
}

export default App;
