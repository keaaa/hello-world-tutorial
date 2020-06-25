import http from 'k6/http';
import { sleep } from 'k6';

export default function() {
  http.get('https://server-keaaa-hello-world-dev.playground.radix.equinor.com/');
  sleep(1);
}