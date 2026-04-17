using KeeAuditLib;
using System.Linq;

namespace KeeAuditPlugin
{
    internal static class ValidationResultExtensions
    {
        public static ValidationResultDto ToDto(this ValidationResult validationResult)
        {
            return new ValidationResultDto
            {
                EntryUuid = validationResult.Entry.Uuid.ToString(),
                EntryTitle = validationResult.Entry.Strings.ReadSafe("Title"),
                GroupPath = validationResult.Entry.ParentGroup.GetFullPath(),
                IsCompliant = validationResult.IsCompliant,
                FailedSpecificationNames = validationResult.FailedSpecifications.Select(spec => spec.GetType().Name).ToArray()
            };
        }
    }
}
