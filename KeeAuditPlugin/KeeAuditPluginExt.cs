using KeeAuditPlugin.Config;
using KeePass.Plugins;
using KeePassLib.Interfaces;

namespace KeeAuditPlugin
{
    public sealed class KeeAuditPluginExt : Plugin
    {
        private IPluginHost _host = null;
        private AuditService _auditService = null;

        public override bool Initialize(IPluginHost host)
        {
            if (host == null) return false;
            _host = host;

            var rabbitMQConfig = new RabbitMQConnectionConfig(host.CustomConfig);
            IStatusLogger logger = _host.MainWindow.CreateShowWarningsLogger();
            try
            {
                logger.StartLogging("KeeAuditPlugin", true);
                _auditService = new AuditService(rabbitMQConfig, _host, logger);
            }
            catch (FailedToConnectException ex)
            {
                logger.SetText(ex.Message, KeePassLib.Interfaces.LogStatusType.Error);
                return false;
            }
            catch (FailedToDeclareQueueException ex)
            {
                logger.SetText(ex.Message, KeePassLib.Interfaces.LogStatusType.Error);
                return false;
            }
            finally
            {
                logger.EndLogging();
            }

            return true;
        }

        public override void Terminate()
        {
            _host = null;
        }
    }
}
