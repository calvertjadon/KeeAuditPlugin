namespace KeeAuditPlugin
{
    internal class ValidationResultDto
    {
        public string EntryUuid { get; set; }
        public string EntryTitle { get; set; }
        public string GroupPath { get; set; }
        public bool IsCompliant { get; set; }
        public string[] FailedSpecificationNames { get; set; }
    }
}
