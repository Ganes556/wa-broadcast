import './main.css';
import 'htmx.org';
import 'htmx-ext-ws';
import * as countries from './countries_code.json';
import './node_modules/flag-icons/css/flag-icons.min.css';
import parsePhoneNumber from 'libphonenumber-js';
import Alpine from 'alpinejs';
import persist from '@alpinejs/persist';

window.Alpine = Alpine;
window.htmx = htmx;

Alpine.plugin(persist);
const loc = new Intl.Locale(navigator.language);

Alpine.data('countries', () => ({
  data: countries.default,
  selected: countries.default.find(
    (country) => country.code === (loc.region || loc.language.toUpperCase())
  ),
  showContent: false,
  updatePhoneNumber(event, index) {
    const input = event.target;
    let phone  = input.value;
    if(input.value) {
      let formated = parsePhoneNumber(phone, this.selected.code.toUpperCase())
      if(formated?.number) {
        let formatedNum = formated.number.replace("+","");
        this.phone[index] = formatedNum;
      }
      return
    }
    this.phone[index] = "";
  },
  toggleContent() {
    this.showContent = !this.showContent;
  },
}));

Alpine.start();
