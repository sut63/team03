import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn';
import MedicalfileCreate from './components/SaveMed';
import DentalExpense from './components/SaveDenExpen';
import Menu from './components/Menu';
import Patient from './components/SavePatient';
import Appointment from './components/SaveAppoint';
import Dentist from './components/SaveDentist';
import Queue from './components/SaveQueue';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/Menu', Menu);
    router.registerRoute('/SaveMed', MedicalfileCreate);
    router.registerRoute('/SaveDenExpen', DentalExpense);
    router.registerRoute('/SavePatient', Patient);
    router.registerRoute('/SaveAppoint', Appointment);
    router.registerRoute('/SaveDentist', Dentist);
    router.registerRoute('/SaveQueue', Queue);
  }
});
