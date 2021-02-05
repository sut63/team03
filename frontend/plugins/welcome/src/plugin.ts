import { createPlugin } from '@backstage/core';
import SignInDentist from './components/SignInDentist';
import SignInNurse from './components/SignInNurse';
import MedicalfileCreate from './components/SaveMed';
import DentalExpense from './components/SaveDenExpen';
import Menu from './components/Menu';
import Patient from './components/SavePatient';
import Appointment from './components/SaveAppoint';
import Dentist from './components/SaveDentist';
import Queue from './components/SaveQueue';
import Nurse from './components/SaveNurse';
import Welcome from './components/Welcome';
import SearchDenExpen from './components/SearchDenExpen';
import SearchMed from './components/SearchMed';
import SearchAppoint from './components/SearchAppoint';
import SearchQueue from './components/SearchQueue';
import SearchDentist from './components/SearchDentist';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Welcome);
    router.registerRoute('/SignInDentist', SignInDentist);
    router.registerRoute('/SignInNurse', SignInNurse);
    router.registerRoute('/Menu', Menu);
    router.registerRoute('/SaveMed', MedicalfileCreate);
    router.registerRoute('/SaveDenExpen', DentalExpense);
    router.registerRoute('/SavePatient', Patient);
    router.registerRoute('/SaveAppoint', Appointment);
    router.registerRoute('/SaveDentist', Dentist);
    router.registerRoute('/SaveQueue', Queue);
    router.registerRoute('/SaveNurse', Nurse);

    router.registerRoute('/SearchAppoint', SearchAppoint);
    router.registerRoute('/SearchDenExpen', SearchDenExpen);
    router.registerRoute('/SearchMed', SearchMed);
    router.registerRoute('/SearchQueue', SearchQueue);
    router.registerRoute('/SearchDentist', SearchDentist);
  }
});