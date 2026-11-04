"""Sample CDC application script used by the acceptance tests.

Only its content is uploaded and compared; the script is never run.
"""

import logging

log = logging.getLogger(__name__)


class EventHandler:
    """Minimal event-handler pattern for a custom Data Connector script."""

    def process(self, context, event):
        log.info("Received event: %s", event)
        return {"status": "processed"}
