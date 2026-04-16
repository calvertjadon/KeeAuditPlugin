using KeeAuditPlugin.Config;
using KeePass.Plugins;

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
            var logger = _host.MainWindow.CreateShowWarningsLogger();
            logger.StartLogging("KeeAuditPlugin", true);
            logger.SetText("Initializing KeeAuditPlugin...", KeePassLib.Interfaces.LogStatusType.Info);

            var rabbitMQConfig = new RabbitMQConnectionConfig(host.CustomConfig);
            try
            {
                logger.SetText("Connecting to RabbitMQ...", KeePassLib.Interfaces.LogStatusType.Info);
                _auditService = new AuditService(rabbitMQConfig);
                logger.SetText("KeeAuditPlugin initialized successfully.", KeePassLib.Interfaces.LogStatusType.Info);
            }
            catch (FailedToConnectException ex)
            {
                logger.SetText(ex.Message, KeePassLib.Interfaces.LogStatusType.Error);
                return false;
            }
            catch (FailedToCreateChannelException ex)
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
